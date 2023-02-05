import { json } from '@sveltejs/kit'

export const prerender = true;

export const GET = async ({ params }) => {
  try {
    const plant = await import(`../../../lib/data/plants/${params.slug}.json`)
    return json(plant)
  } catch (e: any) {
    return json({})
  }
}