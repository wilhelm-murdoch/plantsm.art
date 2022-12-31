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

export interface Badge {
  foreground: string;
  background: string;
  emoji: string;
}

export function isDeadly(plant: CardPlant) {
  return plant.symptoms.filter(s => s.slug == "death").length >= 1
}

export const animalToBadge: { [char: string]: Badge } = {
  cat: {
    foreground: 'yellow-900',
    background: 'yellow-100',
    emoji: '🐱'
  },
  dog: {
    foreground: 'orange-900',
    background: 'orange-100',
    emoji: '🐶'
  },
  horse: {
    foreground: 'green-900',
    background: 'green-100',
    emoji: '🐴'
  },
  rat: {
    foreground: 'slate-900',
    background: 'slate-100',
    emoji: '🐭'
  },
  bird: {
    foreground: 'blue-900',
    background: 'blue-100',
    emoji: '🐦'
  },
  rabbit: {
    foreground: 'slate-900',
    background: 'slate-100',
    emoji: '🐰'
  },
  reptile: {
    foreground: 'green-700',
    background: 'green-50',
    emoji: '🦎'
  },
  hamster: {
    foreground: 'yellow-700',
    background: 'yellow-50',
    emoji: '🐹'
  }
};