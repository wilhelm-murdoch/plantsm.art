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

export type Classification = {
  kingdom: string;
  clades: string[];
  order: string;
  family: string;
  genus: string;
  species: string;
};

export type Plant = {
  pid: string;
  name: string;
  animals: string[];
  common: Slugged[];
  symptoms: Slugged[];
  images: Image[];
  classification: Classification;
  wikipedia_url: string;
  date_last_updated: string;
  index?: string;
};

export type PlantsWrapped = {
  plants: Plant[]
}