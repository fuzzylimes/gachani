<script>
  let links = [
    {
      type: "dropdown",
      url: "#/ranking",
      text: "Rankings",
      items: [
        {
          url: "#/ranking/all",
          text: "All Anime",
        },
        {
          url: "#/ranking/tv",
          text: "Anime TV Shows",
        },
        {
          url: "#/ranking/movie",
          text: "Anime Movies",
        },
        {
          url: "#/ranking/bypopularity",
          text: "Anime By Popularity",
        },
        {
          url: "#/ranking/airing",
          text: "Now Airing",
        },
      ],
    },
    {
      url: "#/Search",
      text: "Search",
    },
    {
      url: "#/about",
      text: "About",
    },
  ];

  // Handle the menu icon
  document.addEventListener("DOMContentLoaded", () => {
    // Get all "navbar-burger" elements
    const burger = document.querySelector(".navbar-burger");
    const links = document.querySelector(".navbar-menu");
    const items = document.querySelectorAll(".navbar-item");
    burger.addEventListener("click", () => {
      links.classList.toggle("is-active");
    });
    items.forEach((i) => {
      i.addEventListener("click", () => {
        if (links.classList.contains("is-active")) {
          links.classList.toggle("is-active");
        }
      });
    });
  });

  // Handle the drop down
  const toggleMenu = (e) => {
    let menu = e.currentTarget.querySelector(".navbar-dropdown");
    if (e.target.parentElement.classList.contains("navbar-dropdown"))
      menu.style.display = "none";
    setTimeout(() => {
      menu.style.display = "";
      e.target.blur();
    }, 100);
  };
</script>

<nav class="navbar is-fixed-top" aria-label="main navigation">
  <div class="container">
    <div class="navbar-brand">
      <a class="navbar-item is-size-4" href="#/">
        <img src="logo.png" alt="GachAni" height="50" />
      </a>

      <!-- svelte-ignore a11y-missing-attribute -->
      <a
        role="button"
        class="navbar-burger"
        aria-label="menu"
        aria-expanded="false"
        data-target="navbarBasicExample"
      >
        <span aria-hidden="true" />
        <span aria-hidden="true" />
        <span aria-hidden="true" />
      </a>
    </div>

    <div id="navbarBasicExample" class="navbar-menu">
      <div class="navbar-end">
        {#each links as link}
          {#if link.type == "dropdown"}
            <div
              class="navbar-item has-dropdown is-hoverable"
              on:click={(event) => toggleMenu(event)}
            >
              <!-- svelte-ignore a11y-missing-attribute -->
              <a class="navbar-link">{link.text}</a>
              <div class="navbar-dropdown">
                {#each link.items as item}
                  <a href={item.url} class="navbar-item">{item.text}</a>
                {/each}
              </div>
            </div>
          {:else}
            <a class="navbar-item" href={link.url}>
              {link.text}
            </a>
          {/if}
        {/each}
      </div>
    </div>
  </div>
</nav>
