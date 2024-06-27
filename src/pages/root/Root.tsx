import Hero from "../../components/hero/Hero.tsx";

export default function Root() {
    return (
        <>
            <Hero image={"background.png"}
                  title={"Architects Land"}
                  description={"Est une famille de SMP Minecraft privé mélangeant semi-RP, survie et mods uniques."}
            />
            <main>
                <div class="content presentation presentation__smp large-text">
                    <h2>SMP ?</h2>
                    <p>
                        Holycube, QSMP et Content SMP appartiennent à un unique type de serveur minecraft unique : les
                        SMP.
                    </p>
                    <p>
                        Un SMP est une simple survie multijoueur.
                    </p>
                    <p>
                        Sauf qu'après avoir fait des centaines de survie, Minecraft ne se renouvelle plus.
                    </p>
                    <p>
                        Nous avons donc décidé de créer des personnages et une histoire pour rendre notre monde plus
                        vivant.
                    </p>
                    <p>
                        Mais même avec ce semi-RP, on arrive à s'ennuyer.
                    </p>
                    <p>
                        On a donc commencé à modifier le jeu de base pour le redécouvrir.
                    </p>
                </div>
                <div class="content large-text presentation presentation__architects-land">
                    <h2>Architects Land</h2>
                    <p>
                        Architects Land a été créé avec cet objectif en tête : refaire vivre Minecraft, même après des
                        milliers d'heures de jeu.
                    </p>
                    <p>
                        Pour le rendre de nouveau intéressant, on doit créer de nouvelles intéractions entre les
                        joueurs, rajouter de la difficulté et recréer l'émerveillement des premiers jours.
                    </p>
                    <p>
                        Chaque saison d'Architects Land a été imaginée pour remettre de la difficulté dans le jeu, pour
                        provoquer l'émerveillement et pour construire un lieu favorable aux intéractions.
                    </p>
                </div>
            </main>
        </>
    )
}
