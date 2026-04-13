package academy.cli;

import academy.cli.converters.AffineParamsConverter;
import academy.cli.converters.FunctionsConverter;
import academy.model.AffineTransformation;
import academy.model.functions.Function;
import java.io.File;
import java.nio.file.Path;
import java.util.List;
import picocli.CommandLine;

public class Options {
    @CommandLine.Option(
            names = {"-w", "--width"},
            description = "Output picture width",
            defaultValue = "1920")
    public int width;

    @CommandLine.Option(
            names = {"-h", "--height"},
            description = "Output picture height",
            defaultValue = "1080")
    public int height;

    @CommandLine.Option(
            names = {"--seed"},
            description = "Generator begin value",
            defaultValue = "5")
    public long seed;

    @CommandLine.Option(
            names = {"-i", "--iteration-count"},
            description = "Generator iterations count",
            defaultValue = "2500")
    public int iterationCount;

    @CommandLine.Option(
            names = {"-o", "--output-path"},
            description = "Output png picture file path",
            defaultValue = "result.png")
    public Path outputPath;

    @CommandLine.Option(
            names = {"-t", "--threads"},
            description = "Threads count",
            defaultValue = "1")
    public int threadsCount;

    @CommandLine.Option(
            names = {"-g", "--gamma-correction"},
            description = "Enables gamma correction")
    public boolean gammaCorrection;

    @CommandLine.Option(
            names = {"--gamma"},
            description = "Gamma value for correction",
            defaultValue = "2.2")
    public double gamma;

    @CommandLine.Option(
            names = {"-s", "--symmetry-level"},
            description = "Number of points turns",
            defaultValue = "1")
    public int symmetryLevel;

    @CommandLine.Option(
            names = {"-ap", "--affine-params"},
            description = "Affine transformations configuration",
            converter = AffineParamsConverter.class,
            defaultValue = "-0.593,-0.851,0.558,-1.229,-0.157,-0.988")
    public List<AffineTransformation.Params> affineTransformations;

    @CommandLine.Option(
            names = {"-f", "--functions"},
            description = "Applied transformation methods configuration",
            converter = FunctionsConverter.class,
            defaultValue = "spherical:1")
    public List<Function> functions;

    @CommandLine.Option(
            names = {"-c", "--config"},
            description = "Path to JSON config file")
    public File configPath;
}
