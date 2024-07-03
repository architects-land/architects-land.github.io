import Hero from "../../components/hero/Hero.tsx";

import "./Root.scss";
import { For, Show } from "solid-js";

import seasons from "../../assets/seasons.json";

export default function Root() {
  return (
    <>
      <Hero
        image={"/terre-des-civilisations/background.webp"}
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
            style={`background-image: url("/architects-land.webp")`}
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
            <For each={Object.entries(seasons)}>
              {(s, i) => (
                <>
                  <Show when={i() % 2 == 0}>
                    <Season
                      title={s[1].name}
                      image={s[1].logo}
                      placeholder={`Logo de ${s[1].name}`}
                      href={s[1].id}
                    >
                      <For each={s[1].information.description}>
                        {(item) => <p>{item}</p>}
                      </For>
                    </Season>
                  </Show>
                  <Show when={i() % 2 != 0}>
                    <Season
                      title={s[1].name}
                      image={s[1].logo}
                      placeholder={`Logo de ${s[1].name}`}
                      href={s[1].id}
                      right
                    >
                      <For each={s[1].information.description}>
                        {(item) => <p>{item}</p>}
                      </For>
                    </Season>
                  </Show>
                </>
              )}
            </For>
          </div>
        </div>
      </main>
    </>
  );
}

function Season(props: any) {
  return (
    <div
      class={"presentation__season is-clickable"}
      data-href={`/season/${props.href}`}
    >
      <div class="is-clickable__animation"></div>
      <Show when={!props.right}>
        <div class="presentation--right">
          <h4>{props.title}</h4>
          {props.children}
        </div>
        <img
          src={props.image}
          alt={props.placeholder}
          class="presentation__season__image"
        />
      </Show>
      <Show when={props.right}>
        <div class="is-clickable__animation"></div>
        <img
          src={props.image}
          alt={props.placeholder}
          class="presentation__season__image"
        />
        <div class="presentation--left">
          <h4>{props.title}</h4>
          {props.children}
        </div>
      </Show>
    </div>
  );
}
