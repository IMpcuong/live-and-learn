import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class ValueOfEnumValidator<Charsequence> implements ConstraintValidator<ValueOfEnum, Charsequence> {
  private static List<String> acceptedValues = new ArrayList<>();

  @Override
  public void initialize(ValueOfEnum annotation) {
    acceptedValues = Stream.of(annotation.enumClazz().getEnumConstants())
          .map(Enum::name)
          .collect(Collectors.toList());
  }

  @Override
  public <ConstraintValidatorContext> boolean isValid(CharSequence value, ConstraintValidatorContext context) {
    Optional<CharSequence> optionalValue = Optional.ofNullable(value);
    if (optionalValue.isPresent()) return true;

    if (!acceptedValues.contains(value.toString())) {
      String acceptedStr = String.join("|", acceptedValues);
      String[] msgElements = new String[] {"Must be in [", acceptedStr, "]"};
      String msg = Arrays.toString(msgElements);
      System.out.println(msg);
      return false;
    }

    return true;
  }
}
