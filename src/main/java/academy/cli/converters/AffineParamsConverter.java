package academy.cli.converters;

import academy.model.AffineTransformation;
import picocli.CommandLine;
import java.util.ArrayList;
import java.util.List;

public class AffineParamsConverter implements CommandLine.ITypeConverter<List<AffineTransformation.Params>> {

    @Override
    public List<AffineTransformation.Params> convert(String s) throws Exception {
        try {
            var affineParams = new ArrayList<AffineTransformation.Params>();
            for (var i : s.split("/")) {
                var params = i.split(",");
                double a = Double.parseDouble(params[0]);
                double b = Double.parseDouble(params[1]);
                double c = Double.parseDouble(params[2]);
                double d = Double.parseDouble(params[3]);
                double e = Double.parseDouble(params[4]);
                double f = Double.parseDouble(params[5]);
                affineParams.add(new AffineTransformation.Params(a, b, c, d, e, f));
            }
            return affineParams;
        } catch (ArrayIndexOutOfBoundsException exception) {
            throw new CommandLine.TypeConversionException(exception.getMessage());
        }
    }
}
