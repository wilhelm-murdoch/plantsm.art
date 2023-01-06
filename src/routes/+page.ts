import type { Plant } from '$lib/types/plant';

const baseUrl: string = import.meta.env.DEV ? "http://localhost:5173" : "https://plantsm.art"

export let _isLoading: boolean = true;

export const load = async () => {
  const getPlants = async () => {
    const res = await fetch(baseUrl + '/plants.json')
    const data = await res.json()

    return data.map((plant: Plant) => ({
      ...plant,
      index: `${plant.name.toLocaleLowerCase()} ${plant.common?.map(c => c.name.toLocaleLowerCase()).join(' ')} ${plant.animals.join(' ')} ${plant.symptoms.map(s => s.name.toLocaleLowerCase()).join(' ')}`
    }));
  }

  _isLoading = false

  return {
    plants: getPlants()
  }
}
