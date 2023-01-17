
interface Animal {
  foreground: string;
  background: string;
  emoji: string;
}

const animals: { [char: string]: Animal } = {
  // text-yellow-800 bg-yellow-100
  cats: {
    foreground: 'yellow-800',
    background: 'yellow-100',
    emoji: '🐱'
  },

  // text-orange-800 bg-orange-100
  dogs: {
    foreground: 'orange-800',
    background: 'orange-100',
    emoji: '🐶'
  },

  // text-green-800 bg-green-100
  horses: {
    foreground: 'green-800',
    background: 'green-100',
    emoji: '🐴'
  },

  // text-slate-800 bg-slate-100
  rats: {
    foreground: 'slate-800',
    background: 'slate-100',
    emoji: '🐭'
  },

  // text-blue-800 bg-blue-100
  birds: {
    foreground: 'blue-800',
    background: 'blue-100',
    emoji: '🐦'
  },

  // text-slate-800 bg-slate-100
  rabbits: {
    foreground: 'slate-800',
    background: 'slate-100',
    emoji: '🐰'
  },

  // text-green-800 bg-green-50
  reptiles: {
    foreground: 'green-700',
    background: 'green-50',
    emoji: '🦎'
  },

  // text-yellow-700 bg-yellow-100
  hamsters: {
    foreground: 'yellow-700',
    background: 'yellow-100',
    emoji: '🐹'
  },

  // text-indigo-700 bg-indigo-100
  fish: {
    foreground: 'indigo-700',
    background: 'indigo-100',
    emoji: '🐟'
  }
};

export const getByAnimal = (animal: string) => {
  var _default = {
    // text-white bg-black
    foreground: 'white',
    background: 'black',
    emoji: '🚫'
  }

  return animal in animals ? animals[animal] : _default;
}

export const getAllAnimals = () => {
  return Object.keys(animals)
}