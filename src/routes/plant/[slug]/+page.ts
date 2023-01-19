export const prerender = true;

// src/routes/blog/[slug]/+page.ts
export async function load(params: any) {
  console.log(params)
  const plant = await import(`../../../lib/data/plants/${params.params.slug}.json`)
  return {
    plant: plant
  }
}