package utils

import "image-processing-service/models"
import "image-processing-service/store_master"

// ConvertStore converts a store_master.Store to models.Store
func ConvertStore(store store_master.Store) models.Store {
    return models.Store{
        StoreID:   store.StoreID,
        StoreName: store.StoreName,
        AreaCode:  store.AreaCode,
        // Any other necessary fields here
    }
}
