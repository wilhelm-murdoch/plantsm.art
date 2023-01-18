import type { PlantSlim } from "$lib/types/plant";

export function isDeadly(plant: PlantSlim) {
  return plant.symptoms.filter(s => s == "death").length >= 1
}