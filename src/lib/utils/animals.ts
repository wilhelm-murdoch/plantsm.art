
interface Animal {
  foreground: string;
  background: string;
  emoji: string;
}

const animals: { [char: string]: Animal } = {
  cats: {
    foreground: 'yellow-900',
    background: 'yellow-100',
    emoji: '🐱'
  },
  dogs: {
    foreground: 'orange-900',
    background: 'orange-100',
    emoji: '🐶'
  },
  horses: {
    foreground: 'green-900',
    background: 'green-100',
    emoji: '🐴'
  },
  rats: {
    foreground: 'slate-900',
    background: 'slate-100',
    emoji: '🐭'
  },
  birds: {
    foreground: 'blue-900',
    background: 'blue-100',
    emoji: '🐦'
  },
  rabbits: {
    foreground: 'slate-900',
    background: 'slate-100',
    emoji: '🐰'
  },
  reptiles: {
    foreground: 'green-700',
    background: 'green-50',
    emoji: '🦎'
  },
  hamsters: {
    foreground: 'yellow-700',
    background: 'yellow-50',
    emoji: '🐹'
  }
};

export const getByAnimal = (animal: string) => {
  var _default = {
    foreground: 'white',
    background: 'black',
    emoji: '🚫'
  }

  return animal in animals ? animals[animal] : _default;
}

export const getAllAnimals = () => {
  return Object.keys(animals)
}