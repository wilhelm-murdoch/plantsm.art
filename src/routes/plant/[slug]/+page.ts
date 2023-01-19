export const prerender = true;

// src/routes/blog/[slug]/+page.ts
export async function load({params}) {
  const plant = await import(`../../../lib/data/plants/${params.slug}.json`)
  return {
    plant: plant
  }
}