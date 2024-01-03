export type Slugged = {
  name: string;
  slug: string;
};

export type Image = {
  source_url: string;
  attribution: string;
  license: string;
  relative_path: string;
};

export type Plant = {
  pid: string;
  name: string;
  animals: string[];
  common: Slugged[];
  symptoms: Slugged[];
  images: Image[];
  wikipedia_url: string;
  date_last_updated: string;
  index?: string;
};

export type PlantSlim = {
  pid: string;
  name: string;
  common: string[];
  symptoms: string[];
  animals: string[];
  common_total: number;
  symptoms_total: number;
  image_total: number;
  cover_image_url: string;
  is_deadly: boolean;
  search_index?: string;
};

export type PlantsWrapped = {
  plants: PlantSlim[]
}

export function isDeadly(plant: PlantSlim) {
  return plant.symptoms.filter((s: any) => s.slug == "death").length >= 1
}