import static java.lang.annotation.ElementType.ANNOTATION_TYPE;
import static java.lang.annotation.ElementType.CONSTRUCTOR;
import static java.lang.annotation.ElementType.FIELD;
import static java.lang.annotation.ElementType.METHOD;
import static java.lang.annotation.ElementType.PARAMETER;
import static java.lang.annotation.ElementType.TYPE_USE;
import static java.lang.annotation.RetentionPolicy.RUNTIME;

import java.lang.annotation.Documented;
import java.lang.annotation.Retention;
import java.lang.annotation.Target;

// import javax.validation.Constraint;
// import javax.validation.Payload;

@Target({METHOD, FIELD, ANNOTATION_TYPE, CONSTRUCTOR, PARAMETER, TYPE_USE})
@Retention(RUNTIME)
@Documented
// @Constraint(validateBy = ValueOfEnumValidator.class)
public @interface ValueOfEnum {
  Class<? extends Enum<?>> enumClazz();

  String message() default "must be any of enum {enumClass}";

  Class<?>[] groups() default {};

  // Example: because I don't want to use any package-mangement's import suggestion.
  // Class<? extends Payload>[] payload() default {};
}
