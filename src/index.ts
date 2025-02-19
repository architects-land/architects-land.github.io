function setupEvents() {
    if (document.querySelector(".not-found") !== null) {
        setTimeout((_: any) => {
            location.replace(`${location.origin}`);
        }, 5 * 1000);
    }

    const nav = document.getElementById("navbar");
    if (nav === null) {
        return;
    }

    if (window.innerWidth > 620) {
        document.addEventListener("scroll", (_) => {
            if (window.scrollY >= window.innerHeight) {
                nav.style.display = "";
                nav.classList.add("nav--present");
                return;
            }
            nav.classList.remove("nav--present");
        });
    }
}

setupEvents();
