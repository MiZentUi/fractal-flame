package academy;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import picocli.CommandLine;

public class BenchmarkTest {
    private static CommandLine app;
    private static long singleDuration;

    @BeforeAll
    static void singleThread() {
        // Arrange
        app = new CommandLine(new Application());

        // Act
        long begin = System.currentTimeMillis();
        app.execute();
        singleDuration = System.currentTimeMillis() - begin;
    }

    @Test
    void threads_2() {
        // Act
        long begin = System.currentTimeMillis();
        app.execute("-t", "2");
        long end = System.currentTimeMillis();

        // Assert
        assertThat(end - begin).isLessThan(singleDuration);
    }

    @Test
    void threads_4() {
        // Act
        long begin = System.currentTimeMillis();
        app.execute("-t", "4");
        long end = System.currentTimeMillis();

        // Assert
        assertThat(end - begin).isLessThan(singleDuration);
    }

//    @Test
//    void threads_8() {
//        // Act
//        long begin = System.currentTimeMillis();
//        app.execute("-t", "8");
//        long end = System.currentTimeMillis();
//
//        // Assert
//        assertThat(end - begin).isLessThan(singleDuration);
//    }
}
