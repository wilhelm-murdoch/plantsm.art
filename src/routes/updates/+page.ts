export const prerender = true;

export const load = async () => {
  const getEvents = async () => {
    const res = await fetch('https://api.github.com/repos/wilhelm-murdoch/plantsm.art/events?per_page=100')
    const data: any[] = await res.json()

    return data.reduce((e, { type, created_at, actor, payload }) => {
      e[actor.display_login as keyof typeof e] = e[actor.display_login as keyof typeof e] || {
        avatar: actor.avatar_url,
        name: actor.display_login,
        url: actor.url.replace('api.', '').replace('users/', ''),
        events: []
      };
      (e[actor.display_login as keyof typeof e] as any).events.push({ type, created_at, payload });
      return e;
    }, {})
  }

  return {
    events: await getEvents()
  }
}