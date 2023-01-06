export interface PlantImage {
  relative_path: string;
}

export interface Term {
  slug: string;
  name: string;
}

export interface CardPlant {
  animals: string[];
  common: Term[];
  pid: string;
  name: string;
  images: PlantImage[];
  symptoms: Term[];
}

export function isDeadly(plant: CardPlant) {
  return plant.symptoms.filter(s => s.slug == "death").length >= 1
}