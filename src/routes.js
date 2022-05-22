// Components
import Home from './routes/Home.svelte'
import About from './routes/About.svelte'
import Ranking from './routes/Ranking.svelte'
import Recommendations from './routes/Recommended.svelte'
import Details from './routes/Details.svelte'
import Search from './routes/Search.svelte'
import NotFound from './routes/NotFound.svelte'

// Export the route definition object
export default {
    '/': Home,
    '/about': About,
    '/ranking/:mode': Ranking,
    '/recommendations/:id': Recommendations,
    '/details/:id': Details,
    '/search': Search,
    '*': NotFound
}