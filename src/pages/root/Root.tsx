import Hero from "../../components/hero/Hero.tsx";

import "./Root.scss";

export default function Root() {
  return (
    <>
      <Hero
        image={"/background.png"}
        title={"Architects Land"}
        description={"Famille de SMP Minecraft privé"}
      />
      <main>
        <div class="content presentation presentation__smp large-text">
          <h2>SMP ?</h2>
          <p>
            Holycube, QSMP et Content SMP appartiennent à un unique type de
            serveur minecraft unique : les SMP.
          </p>
          <p>Un SMP est une simple survie multijoueur.</p>
          <p>
            Sauf qu'après avoir fait des centaines de survie, Minecraft ne se
            renouvelle plus.
          </p>
          <p>
            Nous avons donc décidé de créer des personnages et une histoire pour
            rendre notre monde plus vivant.
          </p>
          <p>Mais même avec ce semi-RP, on arrive à s'ennuyer.</p>
          <p>
            On a donc commencé à modifier le jeu de base pour le redécouvrir.
          </p>
        </div>
        <div class="content large-text presentation presentation__architects-land">
          <div
            class="presentation__architects-land__image"
            style={`background-image: url("/architects-land.png")`}
          ></div>
          <div class="presentation__architects-land__text">
            <h2>Nouvelle forme de SMP</h2>
            <p>
              Architects Land a été créé avec cet objectif en tête : refaire
              vivre Minecraft, même après des milliers d'heures de jeu.
            </p>
            <p>
              Pour le rendre de nouveau intéressant, on doit créer de nouvelles
              intéractions entre les joueurs, rajouter de la difficulté et
              recréer l'émerveillement des premiers jours.
            </p>
            <p>
              Chaque saison d'Architects Land a été imaginée pour remettre de la
              difficulté dans le jeu, pour provoquer l'émerveillement et pour
              construire un lieu favorable aux intéractions.
            </p>
          </div>
        </div>
        <div
          class="content large-text presentation presentation__seasons"
          id={"seasons"}
        >
          <h2>Saisons</h2>
          <div class="presentation__seasons__grid">
            <Season
              title={"Terre des Civilisations"}
              image={"/terre-des-civilisations/logo.png"}
              placeholder={"Logo de Terre des Civilisations"}
              href={"terre-des-civilisations"}
            >
              <p>
                Première saison d'Architects Land, Terre des Civilisations se
                déroule durant l'été 2024, en 1.21.
              </p>
              <p>
                Durant cette saison, la génération fut grandement modifiée pour
                qu'elle devienne plus réaliste. Les villages furent aussi
                améliorer pour qu'ils deviennent plus grand et plus important.
              </p>
              <p>
                Comme les villages occupent une place beaucoup centrale, les
                bases des joueurs doivent forcément être reliées au village le
                plus proche. Cette connection de base forme les clans des
                joueurs.
              </p>
              <p>
                La difficulté du monde est adaptée en temps réelle au niveau des
                joueurs : plus ils meurent, plus le niveau de difficulté
                augmentent jusqu'à atteindre l'ultra hardcore. La map était
                aussi réduite au début pour concentrer un maximum les bases
                autour du 0 0.
              </p>
            </Season>
          </div>
        </div>
      </main>
    </>
  );
}

function Season(props: any) {
  if (props.right) {
    return (
      <div class={"presentation__season is-clickable"} data-href={props.href}>
        <div class="presentation__season__animation"></div>
        <img
          src={props.image}
          alt={props.placeholder}
          class="presentation__season__image"
        />
        <div class="presentation--left">
          <h4>{props.title}</h4>
          {props.children}
        </div>
      </div>
    );
  }
  return (
    <div class={"presentation__season is-clickable"} data-href={props.href}>
      <div class="presentation__season__animation"></div>
      <div class="presentation--right">
        <h4>{props.title}</h4>
        {props.children}
      </div>
      <img
        src={props.image}
        alt={props.placeholder}
        class="presentation__season__image"
      />
    </div>
  );
}
