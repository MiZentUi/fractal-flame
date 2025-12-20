package academy;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.assertj.core.api.AssertionsForClassTypes.fail;

import academy.cli.converters.FunctionsConverter;
import academy.model.functions.PolarFunction;
import academy.model.functions.SphericalFunction;
import academy.model.functions.SwirlFunction;
import java.util.List;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import picocli.CommandLine;

public class FunctionsConverterTest {
    private static FunctionsConverter converter;

    @BeforeAll
    static void init() {
        converter = new FunctionsConverter();
    }

    @Test
    void zero() throws Exception {
        // Arrange
        var functionsString = "swirl:0,polar:0";
        var expected = List.of(new SwirlFunction(0.0), new PolarFunction(0.0));

        // Act
        var functions = converter.convert(functionsString);

        // Assert
        assertThat(functions).isEqualTo(expected);
    }

    @Test
    void floating() throws Exception {
        // Arrange
        var functionsString = "polar:0.1,swirl:0.2,spherical:0.3";
        var expected = List.of(new PolarFunction(0.1), new SwirlFunction(0.2), new SphericalFunction(0.3));

        // Act
        var functions = converter.convert(functionsString);

        // Assert
        assertThat(functions).isEqualTo(expected);
    }

    @Test
    void exception() throws Exception {
        // Arrange
        var functionsString = "swirl:";

        // Act
        try {
            converter.convert(functionsString);
        } catch (CommandLine.TypeConversionException exception) {
            return;
        }

        // Assert
        fail();
    }
}
