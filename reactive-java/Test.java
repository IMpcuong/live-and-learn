import java.util.Arrays;
import java.util.Queue;
import java.util.concurrent.LinkedBlockingDeque;
import java.util.function.BiFunction;
import java.util.function.Consumer;
import java.util.stream.Collectors;

// Source: https://tvd12.com/observer-design-pattern-catch-exception/
public class Test {
  // Example:
  public static interface Callback extends Consumer<Integer> {
    // void onResult(Integer result);

    // void onError(Throwable e);
  }

  public static void main(String[] args) throws Exception {
    BiFunction<Integer, Integer, Integer> divideFuncFail = (dividend, divisor) -> Math.multiplyExact(dividend, 1 / divisor);
    System.out.println(divideFuncFail.apply(100, 10));

    BiFunction<Integer, Integer, Double> divideFuncSuccess = (dividend, divisor) -> Double.valueOf(dividend / divisor);
    System.out.println(divideFuncSuccess.apply(100, 10));

    String[] arrStr = new String[] {"0", "1", "2"};
    var arrToStr = Arrays.asList(arrStr).stream()
          .map(e -> new StringBuilder().append(e).toString())
          .collect(Collectors.joining(""));
    System.out.println(arrToStr);

    var test = new Test();
    var divider = test.new ObserverDivider();
    divider.divide(10, 5, result -> {
      // NOTE: This arithmetric expression was able to execute normally.

      // Example:
      // @Override
      // public void onResult(Integer result) {
      //   System.out.println("10/5 = " + result);
      // }

      // @Override
      // public void onError(Throwable err) {
      //   System.out.println("10/5: " + err);
      // }
      System.out.println("Result ~ 10/5 = " + result);
    });

    Thread.sleep(50);

    divider.divide(10, 0, result -> {
      // NOTE: This arithmetric expression made our applcation freezed.
      System.out.println("Result ~ 10/0 = " + result);
    });

    divider.divide(20, 5, result -> {
      // NOTE: This arithmetric expression was not able to execute normally.
      System.out.println("Result ~ 20/5 = " + result);
    });
  }

  public class ObserverDivider {
    private Queue<DivideReq> _queue = new LinkedBlockingDeque<>();

    private ObserverDivider() {
      Thread _nThread = new Thread(() -> {
        while (true) {
          DivideReq req = _queue.poll();
          if (req != null) {
            try {
              req.callback.accept(req.dividend / req.divisor);
            } catch(Throwable err) {
              err.printStackTrace();
            }
          }
        }
      });
      _nThread.start();
    }

    public void divide(Integer dividend, Integer divisor, Callback callback) {
      _queue.offer(new DivideReq(dividend, divisor, callback));
    }
  }

  public class DivideReq {
    private Integer dividend;
    private Integer divisor;
    private Callback callback;

    public DivideReq(Integer a, Integer b, Callback callback) {
      this.dividend = a;
      this.divisor = b;
      this.callback = callback;
    }
  }
}
