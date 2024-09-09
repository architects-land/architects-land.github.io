import "./Button.scss";

export type Properties = {
  href: string;
  content: string;
};

export default function Button(props: Properties) {
  return (
    <a class={"btn"} href={props.href}>
      {props.content}
    </a>
  );
}
