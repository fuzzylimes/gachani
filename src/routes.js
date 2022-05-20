// Components
import Home from './routes/Home.svelte'
import About from './routes/About.svelte'
import Ranking from './routes/Ranking.svelte'
import NotFound from './routes/NotFound.svelte'

// Export the route definition object
export default {
    '/': Home,
    '/about': About,
    '/ranking/all': Ranking,
    '/ranking/tv': Ranking,
    '/ranking/movie': Ranking,
    '/ranking/bypopularity': Ranking,
    '/ranking/airing': Ranking,
    // '/season': Season,
    // '/related': Related,
    '*': NotFound
}