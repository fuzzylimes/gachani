<script>
    import { onMount } from "svelte";
    import { inview } from "svelte-inview";
    import Card from "../components/Card.svelte";
    import { location } from "svelte-spa-router";

    const mode = $location.split("/")[2];
    let bannerText;
    switch (mode) {
        case "tv":
        case "all":
            bannerText = `Currently viewing ${
                mode[0].toUpperCase() + mode.slice(1).toLowerCase()
            } Anime`;
            break;
        case "movie":
            bannerText = `Currently viewing Anime Movies`;
            break;
        case "bypopularity":
            bannerText = `Currently viewing Anime by user Popularity`;
            break;
        default:
            bannerText = `Currently airing Anime`;
            break;
    }
    let next = "";
    let hasNext = false;
    let results = [];
    let newResults = [];

    $: results = [...results, ...newResults];

    const fetchData = async () => {
        const res = await fetch(`/api/ranking?ranking_type=${mode}&${next}`);
        const resJson = await res.json();
        newResults = resJson.data.map((n) => n.node);
        if (resJson?.paging?.next) {
            const nextLink = resJson.paging.next;
            next = nextLink.split("?").slice(1).join("?");
            hasNext = true;
        } else {
            hasNext = false;
        }
        console.log(resJson);
    };

    onMount(() => {
        fetchData();
    });

    const handleChange = (e) => {
        if (e.detail.inView && hasNext) fetchData();
    };
</script>

<main>
    <div class="container px-5">
        <div
            class="columns has-text-centered is-centered is-multiline is-variable is-10-mobile is-two-thirds-tablet is-half-desktop is-one-third-widescreen"
        >
            <div class="column is-full">
                <div class="section is-medium">
                    <h1 class="title is-1">Anime Rankings</h1>
                    <h2 class="subtitle is-4">
                        {bannerText}
                    </h2>
                </div>
            </div>
            {#each results as result}
                <div class="column is-5-tablet is-4-desktop">
                    <Card anime={result} />
                </div>
            {/each}
            <!-- This component is place RIGHT BELOW the last post fetched. handleChange will be triggered
            as soon as it appears on frame. We'll use this info to know when to fetch more data  -->
            <div use:inview={{}} on:change={handleChange} />
        </div>
    </div>
</main>

<style>
</style>
