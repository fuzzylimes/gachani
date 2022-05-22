<script>
    import Card from "../components/Card.svelte";
    import Spinner from "../components/Spinner.svelte";

    export let params = {};

    let id = params.id;
    let results = [];
    let newResults = [];
    let title = "";
    let isLoading = true;

    $: results = [...results, ...newResults];
    $: params, remount();

    const fetchData = async () => {
        const res = await fetch(`/api/recommended?id=${id}`);
        const resJson = await res.json();
        title = resJson.title;
        newResults = resJson.recommendations.map((n) => n.node);
        isLoading = false;
    };

    const remount = () => {
        id = params.id;
        results = [];
        newResults = [];
        isLoading = true;
        fetchData();
    };
</script>

<main>
    <div class="container px-5">
        <div
            class="columns has-text-centered is-centered is-multiline is-variable is-10-mobile is-two-thirds-tablet is-half-desktop is-one-third-widescreen"
        >
            <div class="column is-full">
                <div class="section is-medium">
                    <h1 class="title is-1">Anime Recommendations</h1>
                    <h2 class="subtitle is-4">
                        {#if title != "" && !isLoading}
                            If you like {title}, you might also like...
                        {/if}
                    </h2>
                </div>
            </div>
            {#if isLoading}
                <Spinner />
            {/if}
            {#each results as result}
                <div class="column is-5-tablet is-4-desktop">
                    <Card anime={result} />
                </div>
            {/each}
        </div>
    </div>
</main>

<style>
</style>
