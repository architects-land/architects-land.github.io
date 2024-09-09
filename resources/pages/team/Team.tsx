import Hero from "../../components/hero/Hero.tsx";
import "./Team.scss";
import List from "../../components/people/List.tsx";
import Card from "../../components/people/Card.tsx";

export default function Team() {
  return (
    <>
      <Hero
        image={"/village-night.webp"}
        title={"Équipe"}
        description={""}
        min
      />
      <main>
        <div class="content team__presentation large-text">
          <h2>Les créateurs d'Architects Land</h2>
          <p>
            Architects Land, c'est aussi une équipe dévouée visant à créer un
            SMP inoubliable !
          </p>
          <p>
            Architects Land n'est pas qu'un serveur Minecraft, c'est aussi des
            modpacks avec des mods customs, des événements nécessitant des
            plugins créés de toute pièce et une site web pour présenter le
            projet.
          </p>
        </div>
        <List>
          <Card
            name={"Anhgelus Morhtuuzh"}
            image={"/skins/anhgelus.png"}
            description={
              "Fondateur et développeur, Anhgelus Morhtuuzh gère le serveur et prend toutes les décisions importantes."
            }
            link={"https://youtube.com/@anhgelus"}
          />
          <Card
            name={"Akik4"}
            image={"/skins/Akik4.png"}
            description={
              "Akik4 est un développeur Spigot aidant grandement Anhgelus Morhtuuzh."
            }
            link={"https://github.com/Akik4"}
          />
          <Card
            name={"Léo-21"}
            image={"/skins/Leo_21_.png"}
            description={"Léo-21 développe le mod Sneaky Player Names."}
            link={"https://github.com/leo-210"}
          />
        </List>
      </main>
    </>
  );
}
