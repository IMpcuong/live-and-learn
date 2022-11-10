public interface ConstraintValidator<T1, T2> {
  public void initialize(ValueOfEnum annotation);

  public <ConstraintValidatorContext> boolean isValid(CharSequence value, ConstraintValidatorContext context);
}
