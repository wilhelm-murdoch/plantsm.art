export const prerender = true;

// export const load = async () => {
//   const getContributors = async () => {
//     const res = await fetch('https://api.github.com/repos/wilhelm-murdoch/plantsm.art/contributors')
//     const data: any[] = await res.json()

//     return data.map(c => {
//       return {
//         contributions: c.contributions,
//         avatar: c.avatar_url,
//         name: c.login,
//         url: c.html_url
//       }
//     });
//   }

//   return {
//     contributors: getContributors()
//   }
// }