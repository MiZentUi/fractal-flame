package academy;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

import academy.model.Point;
import academy.model.functions.DiscFunction;
import academy.model.functions.EyefishFunction;
import academy.model.functions.HeartFunction;
import academy.model.functions.HorseshoeFunction;
import academy.model.functions.LinearFunction;
import academy.model.functions.PolarFunction;
import academy.model.functions.SinusoidalFunction;
import academy.model.functions.SphericalFunction;
import academy.model.functions.SwirlFunction;
import org.junit.jupiter.api.Test;

public class FunctionTest {
    @Test
    void discTest() {
        // Arrange
        var function = new DiscFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(-0.19259257104834554, -0.24652091454128053);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }

    @Test
    void eyefishTest() {
        // Arrange
        var function = new EyefishFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(1.2880160864200751, 1.9320241296301126);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }

    @Test
    void heartTest() {
        // Arrange
        var function = new HeartFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(5.1921874930473235, -5.00411720856366);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }

    @Test
    void horseshoeTest() {
        // Arrange
        var function = new HorseshoeFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(-2.773500981126146, 6.65640235470275);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }

    @Test
    void linearTest() {
        // Arrange
        var function = new LinearFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(4, 6);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }

    @Test
    void polarTest() {
        // Arrange
        var function = new PolarFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(0.3128329581890012, 6.211102550927978);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }

    @Test
    void sinusoidalTest() {
        // Arrange
        var function = new SinusoidalFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(-0.7568024953079282, -0.27941549819892586);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }

    @Test
    void sphericalTest() {
        // Arrange
        var function = new SphericalFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(0.07692307692307693, 0.11538461538461539);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }

    @Test
    void swirlTest() {
        // Arrange
        var function = new SwirlFunction();
        var point = new Point(4, 6);
        var expectedPoint = new Point(4.495235398082071, -5.638515648273925);

        // Act
        point = function.transform(point);

        // Assert
        assertThat(point).isEqualTo(expectedPoint);
    }
}
