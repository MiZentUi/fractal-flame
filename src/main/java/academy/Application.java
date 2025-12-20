package academy;

import academy.cli.AppConfig;
import academy.cli.Options;
import academy.model.AffineTransformation;
import academy.utils.Generator;
import academy.utils.ImageWriter;
import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectReader;
import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import java.io.IOException;
import java.io.UncheckedIOException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import picocli.CommandLine;
import picocli.CommandLine.Command;

@Command(name = "Fractal Flames", version = "1.0", mixinStandardHelpOptions = true)
public class Application implements Runnable {
    private static final Logger LOGGER = LoggerFactory.getLogger(Application.class);
    private static final ObjectReader JSON_READER = new ObjectMapper(new JsonFactory())
            .setPropertyNamingStrategy(PropertyNamingStrategies.SNAKE_CASE)
            .findAndRegisterModules()
            .reader();

    @CommandLine.Mixin
    private Options options;

    public static void main(String[] args) {
        int exitCode = new CommandLine(new Application()).execute(args);
        System.exit(exitCode);
    }

    @Override
    public void run() {
        var config = loadConfig();
        LOGGER.atInfo().addKeyValue("config", config).log("Config content");

        AffineTransformation.setSeed(config.getSeed());

        var generator = new Generator(config);
        var writer = new ImageWriter(config.getOutputPath());

        generator.render();
        if (config.isGammaCorrection()) {
            generator.correction();
        }
        writer.write(generator.getImage());
    }

    private AppConfig loadConfig() {
        // fill with cli options
        if (options.configPath == null) {
            return new AppConfig(
                    new AppConfig.Size(options.width, options.height),
                    options.iterationCount,
                    options.outputPath,
                    options.threadsCount,
                    options.seed,
                    options.gammaCorrection,
                    options.gamma,
                    options.symmetryLevel,
                    options.functions,
                    options.affineTransformations);
        }

        // use config file if provided
        try {
            return JSON_READER.readValue(options.configPath, AppConfig.class);
        } catch (IOException e) {
            throw new UncheckedIOException(e);
        }
    }
}
