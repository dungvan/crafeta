package main

// Resource return rate threshold - per-mille
var rrrThreshold = 1000

func getTotalCraftedProducts(initResources, numberOfCraftingResourcesPerProduct, productCraftingPrice, rrr int) (totalProducts, totalCraftingCost, resourcesRemaining int) {
	totalProducts = initResources / numberOfCraftingResourcesPerProduct
	for resourcesRemaining := totalProducts*numberOfCraftingResourcesPerProduct*rrr/rrrThreshold + initResources%numberOfCraftingResourcesPerProduct; resourcesRemaining > numberOfCraftingResourcesPerProduct; resourcesRemaining = (resourcesRemaining/numberOfCraftingResourcesPerProduct)*numberOfCraftingResourcesPerProduct*rrr/rrrThreshold + resourcesRemaining%numberOfCraftingResourcesPerProduct {
		totalProducts += resourcesRemaining / numberOfCraftingResourcesPerProduct
	}
	return
}
