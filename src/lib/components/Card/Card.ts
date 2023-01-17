export interface PlantImage {
  relative_path: string;
}

export interface Term {
  slug: string;
  name: string;
}

export interface CardPlant {
  pid: string;
  name: string;
  common: string[];
  common_total: number;
  symptoms: string[];
  symptoms_total: number;
  animals: string[];
  cover_image_url: string;
  image_total: number;
  search_index: string;
  is_deadly: boolean;
}

export function isDeadly(plant: CardPlant) {
  return plant.symptoms.filter(s => s == "death").length >= 1
}