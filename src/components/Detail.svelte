<script>
    import { calcWatchTime, calRunTime } from "../util/time";
    export let anime;

    const scrollToView = (id) => {
        const element = document.getElementById(id);
        if (element) {
            element.scrollIntoView(true); // `true` is for aligning top of element to top of screen
        }
    };
</script>

<div class="card">
    <div class="card-content">
        <div class="column is-full">
            <div class="py-3">
                {#if anime?.alternative_titles?.en}
                    <p class="title is-1">
                        {anime?.alternative_titles?.en}
                    </p>
                    {#if anime.alternative_titles.en !== anime.title}
                        <p class="subtitle is-4">{anime.title}</p>
                    {/if}
                {:else}
                    <p class="title is-1">
                        {anime.title}
                    </p>
                {/if}
            </div>
        </div>
        <div class="columns is-vcentered">
            <div class="column">
                <img
                    src={anime?.main_picture?.large ||
                        anime?.main_picture?.medium}
                    alt="{anime.name}-image"
                />
            </div>
            <div class="column has-text-left">
                <div>
                    <p>
                        <strong>Genres:</strong>
                        {anime?.genres?.map((g) => g.name).join(", ") || "???"}
                    </p>
                    <p>
                        <strong>Media Type:</strong>
                        {anime?.media_type?.toUpperCase() || "???"}
                    </p>
                    <p>
                        <strong>First Aired:</strong>
                        {anime.start_date || "???"}
                    </p>
                    <p>
                        <strong>Last Aired:</strong>
                        {anime.end_date || "???"}
                    </p>
                    <p>
                        <strong>Studios:</strong>
                        {anime?.studios?.map((s) => s.name).join(", ") || "???"}
                    </p>
                    <p>
                        <strong>Rating:</strong>
                        {anime?.rating?.toUpperCase() || "???"}
                    </p>
                    <p>
                        <strong>Episodes:</strong>
                        {anime.num_episodes || "???"}
                    </p>
                    <p>
                        <strong>Run Time:</strong>
                        {calRunTime(anime.average_episode_duration)}
                    </p>
                    <p>
                        <strong>Time to watch:</strong>
                        {calcWatchTime(
                            anime.num_episodes,
                            anime.average_episode_duration
                        )}
                    </p>
                </div>
                <div class="mt-3">
                    <p>
                        <strong>Avg User Score:</strong>
                        {anime.mean || "???"}
                    </p>
                    <p>
                        <strong>MAL Site Ranking:</strong>
                        {anime.rank || "???"}
                    </p>
                    <p>
                        <strong>MAL Popularity Ranking:</strong>
                        {anime.popularity || "???"}
                    </p>
                    <p />
                </div>
                <div class="mt-3">
                    <p>
                        <!-- svelte-ignore a11y-missing-attribute -->
                        <a on:click={() => scrollToView("related")}
                            ><strong>Related</strong></a
                        >
                        |
                        <!-- svelte-ignore a11y-missing-attribute -->
                        <a on:click={() => scrollToView("recommendations")}
                            ><strong>Recommendations</strong></a
                        >
                        |
                        <a
                            href="https://myanimelist.net/anime/{anime.id}"
                            target="_blank"><strong>View on MAL</strong></a
                        >
                    </p>
                </div>
            </div>
        </div>

        <div class="content has-text-left">
            <div class="pxy-2">
                {anime.synopsis}
            </div>
        </div>
    </div>
</div>

<style>
    .card {
        margin-top: 4rem;
    }
</style>
