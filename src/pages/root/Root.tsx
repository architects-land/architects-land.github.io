import Hero from "../../components/hero/Hero.tsx";

export default function Root() {
    return (
        <>
            <Hero image={"background.png"} title={"Architects Land"} description={"Est une famille de SMP Minecraft privé mélengeant semi-RP, survie et mods uniques."} />
            <main>
                <div class="content presentation large-text">
                    <h2>SMP ?</h2>
                    <p>
                        Holycube, QSMP et Content SMP appartiennent à un unique type de serveur minecraft unique : les SMP.
                    </p>
                    <p>
                        Un SMP est une simple survie multijoueur.
                    </p>
                    <p>
                        Sauf qu'après avoir fait des centaines de survie, Minecraft ne se renouvelle plus.
                    </p>
                    <p>
                        Nous avons donc décidé de créer des personnages et une histoire pour rendre notre monde plus vivant.
                    </p>
                    <p>
                        Mais même avec ce semi-RP, on arrive à s'ennuyer.
                    </p>
                    <p>
                        On a donc commencé à modifier le jeu de base pour le redécouvrir.
                    </p>
                </div>
            </main>
        </>
    )
}
