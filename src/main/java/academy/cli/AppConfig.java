package academy.cli;

import academy.model.AffineTransformation;
import academy.model.functions.Function;
import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import java.nio.file.Path;
import java.util.List;

@JsonNaming(PropertyNamingStrategies.SnakeCaseStrategy.class)
public class AppConfig {
    public record Size(int width, int height) {}

    private Size size;
    private int iterationCount;
    private Path outputPath;
    private int threads;
    private long seed;
    private boolean gammaCorrection;
    private double gamma;
    private int symmetryLevel;


    @JsonDeserialize(contentUsing = FunctionDeserializer.class)
    private List<Function> functions;

    private List<AffineTransformation.Params> affineParams;

    public AppConfig() {}

    public AppConfig(Size size, int iterationCount, Path outputPath, int threads, long seed, boolean gammaCorrection, double gamma, int symmetryLevel, List<Function> functions, List<AffineTransformation.Params> affineParams) {
        this.size = size;
        this.iterationCount = iterationCount;
        this.outputPath = outputPath;
        this.threads = threads;
        this.seed = seed;
        this.gammaCorrection = gammaCorrection;
        this.gamma = gamma;
        this.symmetryLevel = symmetryLevel;
        this.functions = functions;
        this.affineParams = affineParams;
    }

    public Size getSize() {
        return size;
    }

    public int getIterationCount() {
        return iterationCount;
    }

    public Path getOutputPath() {
        return outputPath;
    }

    public int getThreads() {
        return threads;
    }

    public long getSeed() {
        return seed;
    }

    public boolean isGammaCorrection() {
        return gammaCorrection;
    }

    public double getGamma() {
        return gamma;
    }

    public int getSymmetryLevel() {
        return symmetryLevel;
    }

    public List<Function> getFunctions() {
        return functions;
    }

    public List<AffineTransformation.Params> getAffineParams() {
        return affineParams;
    }
}
