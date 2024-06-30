import "./People.scss";

export default function List(props: any) {
  return <div class={"content people__list large-text"}>{props.children}</div>;
}
