import { InputHTMLAttributes, ReactNode, useId } from "react";
import styles from "./styles.module.css";

type Props = InputHTMLAttributes<HTMLInputElement> & {
  type?: "text" | "password" | "email" | "number" | "search";
  label?: string;
  error?: string;
  className?: string;
  trailingIcon?: ReactNode;
  leadingIcon?: ReactNode;
};

export default function TextField({
  className,
  id,
  type = "text",
  label,
  leadingIcon,
  trailingIcon,
  error,
  ...restProps
}: Props) {
  const cid = !id ? useId() : id;
  let classes = `${styles.input} ${error ? styles.input_error : ''}`;
  return (
    <div className={classes}>
      {label && <label htmlFor={id || cid}>{label}</label>}
      <div className={`${styles.input__box} ${className ? className : ""}`}>
        {leadingIcon && (
          <span
            className={`${styles.input__icon} ${styles.input__icon__leading}`}
          >
            {leadingIcon}
          </span>
        )}
        <input {...restProps} id={id || cid} />
        {trailingIcon && (
          <span
            className={`${styles.input__icon} ${styles.input__icon__trailing}`}
          >
            {trailingIcon}
          </span>
        )}
      </div>
      {error && <span className={styles.error__message}>{error}</span>}
    </div>
  );
}
