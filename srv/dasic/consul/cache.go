package consul

import (
	"fmt"
	"sync"
	"time"
)

type cachedService struct {
	serviceID   string
	serviceName string
	address     string
	port        int
	tags        []string
	registeredAt time.Time
}

type ServiceRegistry struct {
	client          *Client
	mu              sync.RWMutex
	registeredCache map[string]*cachedService
}

var (
	registry     *ServiceRegistry
	registryOnce sync.Once
)

func NewServiceRegistry(addr string) (*ServiceRegistry, error) {
	client, err := NewClient(addr)
	if err != nil {
		return nil, err
	}

	registryOnce.Do(func() {
		registry = &ServiceRegistry{
			client:          client,
			registeredCache: make(map[string]*cachedService),
		}
	})

	return registry, nil
}

func GetServiceRegistry() *ServiceRegistry {
	return registry
}

func (r *ServiceRegistry) Register(serviceID, serviceName, address string, port int, tags []string) error {
	err := r.client.RegisterService(serviceID, serviceName, address, port, tags)
	if err != nil {
		return fmt.Errorf("failed to register service: %w", err)
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.registeredCache[serviceID] = &cachedService{
		serviceID:   serviceID,
		serviceName: serviceName,
		address:     address,
		port:        port,
		tags:        tags,
		registeredAt: time.Now(),
	}

	return nil
}

func (r *ServiceRegistry) RegisterWithTTL(serviceID, serviceName, address string, port int, tags []string, ttl string) error {
	err := r.client.RegisterServiceWithTTL(serviceID, serviceName, address, port, tags, ttl)
	if err != nil {
		return fmt.Errorf("failed to register service with TTL: %w", err)
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.registeredCache[serviceID] = &cachedService{
		serviceID:   serviceID,
		serviceName: serviceName,
		address:     address,
		port:        port,
		tags:        tags,
		registeredAt: time.Now(),
	}

	return nil
}

func (r *ServiceRegistry) Deregister(serviceID string) error {
	err := r.client.DeregisterService(serviceID)
	if err != nil {
		return fmt.Errorf("failed to deregister service: %w", err)
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.registeredCache, serviceID)

	return nil
}

func (r *ServiceRegistry) GetRegisteredServices() map[string]*cachedService {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make(map[string]*cachedService)
	for k, v := range r.registeredCache {
		result[k] = v
	}

	return result
}

func (r *ServiceRegistry) GetClient() *Client {
	return r.client
}
