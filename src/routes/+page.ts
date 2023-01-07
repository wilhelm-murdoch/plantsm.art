import plants from './plants.json';

export const load = async () => {
  return {
    plants: plants.map((plant: any) => ({
      ...plant,
      index: `${plant.name.toLocaleLowerCase()} ${plant.common?.map((c: any) => c.name.toLocaleLowerCase()).join(' ')} ${plant.animals.join(' ')} ${plant.symptoms.map((s: any) => s.name.toLocaleLowerCase()).join(' ')}`
    }))
  }
}
