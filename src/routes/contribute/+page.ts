export const prerender = true;

export const load = async () => {
  const getContributors = async () => {
    const res = await fetch('https://api.github.com/repos/wilhelm-murdoch/plantsm.art/contributors')
    const data = await res.json()

    return data;
  }

  return {
    contributors: getContributors()
  }
}