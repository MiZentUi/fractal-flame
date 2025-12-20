package academy;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.assertj.core.api.AssertionsForClassTypes.fail;

import academy.cli.converters.AffineParamsConverter;
import academy.model.AffineTransformation;
import java.util.List;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import picocli.CommandLine;

public class AffineParamsConverterTest {
    private static AffineParamsConverter converter;

    @BeforeAll
    static void init() {
        converter = new AffineParamsConverter();
    }

    @Test
    void zero() throws Exception {
        // Arrange
        var affineParamsString = "0,0,0,0,0,0/0,0,0,0,0,0";
        var expected = List.of(
                new AffineTransformation.Params(0, 0, 0, 0, 0, 0), new AffineTransformation.Params(0, 0, 0, 0, 0, 0));

        // Act
        var params = converter.convert(affineParamsString);

        // Assert
        assertThat(params).isEqualTo(expected);
    }

    @Test
    void floating() throws Exception {
        // Arrange
        var affineParamsString = "0.1,0.2,0.3,0.4,0.5,0.6/1.1,1.2,1.3,1.4,1.5,1.6";
        var expected = List.of(
                new AffineTransformation.Params(0.1, 0.2, 0.3, 0.4, 0.5, 0.6),
                new AffineTransformation.Params(1.1, 1.2, 1.3, 1.4, 1.5, 1.6));

        // Act
        var params = converter.convert(affineParamsString);

        // Assert
        assertThat(params).isEqualTo(expected);
    }

    @Test
    void exception() throws Exception {
        // Arrange
        var affineParamsString = "0.1,0.2,0.3,0.4,0.5";

        // Act
        try {
            converter.convert(affineParamsString);
        } catch (CommandLine.TypeConversionException exception) {
            return;
        }

        // Assert
        fail();
    }
}
