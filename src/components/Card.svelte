<script>
    export let anime;

    const calcWatchTime = (episodes, runtime) => {
        if (!episodes || !runtime) {
            return "???";
        }
        const totalSeconds = episodes * runtime;
        const days = Math.floor(totalSeconds / 60 / 60 / 24);
        const hours = Math.floor(totalSeconds / 60 / 60) % 24;
        const minutes = Math.floor(totalSeconds / 60) % 60;
        return `${days} days, ${hours} hours, ${minutes} minutes`;
    };

    const calRunTime = (runtime) => {
        if (!runtime) {
            return "???";
        }

        return `${Math.floor(runtime / 60)} mins`;
    };
</script>

<div class="card">
    <div class="card-image">
        <figure class="image is-3by4">
            <img src={anime?.main_picture?.medium} alt="{anime.name}-image" />
        </figure>
    </div>
    <div class="card-content">
        <div class="media">
            <div class="media-content">
                {#if anime?.alternative_titles?.en}
                    <p class="title is-4 is-link">
                        {anime?.alternative_titles?.en}
                    </p>
                    <p class="subtitle is-6">{anime.title}</p>
                {:else}
                    <p class="title is-4 is-link">
                        {anime.title}
                    </p>
                {/if}
            </div>
        </div>

        <div class="content has-text-left">
            <p>
                <strong>Genres:</strong>
                {anime?.genres?.map((g) => g.name).join(", ") || "???"}
            </p>
            <p><strong>Aired:</strong> {anime.start_date || "???"}</p>
            <p><strong>Episodes:</strong> {anime.num_episodes || "???"}</p>
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
    </div>
    <footer class="card-footer">
        <a href="#/recommendations/{anime.id}" class="card-footer-item"
            >Recommendations</a
        >
        <a href="#/details/{anime.id}" class="card-footer-item">Details</a>
    </footer>
</div>

<style>
    .card {
        background-color: #f2f1ef;
        height: 100%;
        display: flex;
        flex-direction: column;
    }
    .card-footer {
        margin-top: auto;
    }
</style>
