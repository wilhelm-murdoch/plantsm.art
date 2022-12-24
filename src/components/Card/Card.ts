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

export const animalToBadge: { [char: string]: Badge } = {
  cat: {
    foreground: 'text-yellow-900',
    background: 'bg-yellow-100',
    emoji: '🐱'
  },
  dog: {
    foreground: 'text-orange-900',
    background: 'bg-orange-100',
    emoji: '🐶'
  },
  horse: {
    foreground: 'text-green-900',
    background: 'bg-green-100',
    emoji: '🐴'
  },
  rat: {
    foreground: 'text-slate-900',
    background: 'bg-slate-100',
    emoji: '🐭'
  },
  bird: {
    foreground: 'text-blue-900',
    background: 'bg-blue-100',
    emoji: '🐦'
  },
  rabbit: {
    foreground: 'text-gray-900',
    background: 'bg-gray-100',
    emoji: '🐰'
  },
  reptile: {
    foreground: 'text-gray-900',
    background: 'bg-gray-100',
    emoji: '🦎'
  },
  hamster: {
    foreground: 'text-gray-900',
    background: 'bg-gray-100',
    emoji: '🐹'
  }
};