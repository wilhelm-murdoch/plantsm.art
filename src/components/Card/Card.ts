export interface PlantImage {
  url: string;
}

export interface CardPlant {
  animals: string[];
  common: string[];
  id: string;
  name: string;
  images: PlantImage[];
  symptoms: string[];
}

export interface Badge {
  foreground: string;
  background: string;
  emoji: string;
}