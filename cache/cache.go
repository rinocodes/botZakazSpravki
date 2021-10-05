package cache

import (
	"sync"
	"time"
)

var pointertocache *cache

type userData struct {
	organization string
}

type userCache struct {
	userData   userData
	created    time.Time
	expiration int
}

type cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	userCache         map[int]userCache
}

func InitUsersCacheByDefault() *cache {

	newCache := cache{}
	newCache.defaultExpiration = 10 * time.Minute
	newCache.cleanupInterval = 2 * time.Minute

	pointertocache = &newCache

	return &newCache

}

func AddUserCache(defaultExpiration, cleanupInterval time.Duration) {

	// userCache := make(map[string]userCache)

	// cache := Cache{
	// 	items:             items,
	// 	defaultExpiration: defaultExpiration,
	// 	cleanupInterval:   cleanupInterval,
	// }

	// // Если интервал очистки больше 0, запускаем GC (удаление устаревших элементов)
	// if cleanupInterval > 0 {
	// 	cache.StartGarbageCollection() // данный метод рассматривается ниже
	// }

	// return &cache
}

// func (c *Cache) Set(key string, value interface{}) {

// 	var expiration int64

// 	// Если продолжительность жизни равна 0 - используется значение по-умолчанию

// 	duration := c.defaultExpiration

// 	// Устанавливаем время истечения кеша
// 	if duration > 0 {
// 		expiration = time.Now().Add(duration).UnixNano()
// 	}

// 	c.Lock()

// 	defer c.Unlock()

// 	c.items[key] = Item{
// 		Value:      value,
// 		Expiration: expiration,
// 		Created:    time.Now(),
// 	}

// }

// func (c *Cache) Get(key string) interface{} {

// 	c.RLock()

// 	defer c.RUnlock()

// 	item, found := c.items[key]

// 	// ключ не найден
// 	if !found {
// 		return nil
// 	}

// 	// Проверка на установку времени истечения, в противном случае он бессрочный
// 	if item.Expiration > 0 {

// 		// Если в момент запроса кеш устарел возвращаем nil
// 		if time.Now().UnixNano() > item.Expiration {
// 			return nil
// 		}

// 	}

// 	return item.Value
// }

// func (c *Cache) Delete(key string) error {

// 	c.Lock()

// 	defer c.Unlock()

// 	if _, found := c.items[key]; !found {
// 		return errors.New("Key not found")
// 	}

// 	delete(c.items, key)

// 	return nil
// }

// func (c *Cache) StartGarbageCollection() {
// 	go c.GC()
// }

// func (c *Cache) GC() {

// 	for {

// 		<-time.After(c.cleanupInterval)

// 		if c.items == nil {
// 			return
// 		}

// 		// Ищем элементы с истекшим временем жизни и удаляем из хранилища
// 		if keys := c.expiredKeys(); len(keys) != 0 {
// 			c.clearItems(keys)

// 		}

// 	}

// }

// // expiredKeys возвращает список "просроченных" ключей
// func (c *Cache) expiredKeys() (keys []string) {

// 	c.RLock()

// 	defer c.RUnlock()

// 	for k, i := range c.items {
// 		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
// 			keys = append(keys, k)
// 		}
// 	}

// 	return
// }

// // clearItems удаляет ключи из переданного списка, в нашем случае "просроченные"
// func (c *Cache) clearItems(keys []string) {

// 	c.Lock()

// 	defer c.Unlock()

// 	for _, k := range keys {
// 		delete(c.items, k)
// 	}
// }

// // type Item struct {
// //     Value      interface{}
// //     Created    time.Time
// //     Expiration int64
// // }