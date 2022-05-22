<script>
    import Card from "../components/Card.svelte";
    import Detail from "../components/Detail.svelte";
    import Spinner from "../components/Spinner.svelte";

    export let params = {};

    let id = params.id;
    let data;
    let recommendations = [];
    let related = [];
    let title = "";
    let isLoading = true;

    $: params, remount();

    const fetchData = async () => {
        const res = await fetch(`/api/details?id=${id}`);
        const resJson = await res.json();
        data = resJson;
        recommendations = resJson?.recommendations?.map((n) => n.node) || [];
        related =
            resJson?.related_anime?.map((n) => ({
                ...n.node,
                relationship: n.RelationTypeFormatted,
            })) || [];
        isLoading = false;
    };

    const remount = () => {
        related = [];
        recommendations = [];
        id = params.id;
        data = null;
        isLoading = true;
        fetchData();
    };
</script>

<main>
    <div class="container px-5">
        <div
            class="columns has-text-centered is-centered is-multiline is-variable is-10-mobile is-two-thirds-tablet is-half-desktop is-one-third-widescreen"
        >
            {#if isLoading}
                <Spinner />
            {/if}

            <!-- Detail -->
            {#if data}
                <Detail anime={data} />
            {/if}

            <!-- Related -->
            {#if related.length > 0}
                <div class="column is-full" id="related">
                    <div class="section">
                        <h2 class="title is-2">Related Anime</h2>
                    </div>
                </div>
                {#each related as result}
                    <div class="column is-5-tablet is-4-desktop">
                        <Card anime={result} />
                    </div>
                {/each}
            {/if}

            <!-- Recommendations -->
            {#if recommendations.length > 0}
                <div class="column is-full" id="recommendations">
                    <div class="section">
                        <h2 class="title is-2">Recommended Anime</h2>
                    </div>
                </div>
                {#each recommendations as result}
                    <div class="column is-5-tablet is-4-desktop">
                        <Card anime={result} />
                    </div>
                {/each}
            {/if}
        </div>
    </div>
</main>

<style>
</style>
