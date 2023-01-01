
export interface Animal {
  foreground: string;
  background: string;
  emoji: string;
}

export const affectedAnimals: { [char: string]: Animal } = {
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