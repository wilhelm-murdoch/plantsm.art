import plants from '$lib/data/slim.json';

export const load = async () => {
  return {
    plants: plants
  }
}
