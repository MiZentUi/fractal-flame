package academy.cli.converters;

import academy.model.functions.Function;
import academy.utils.FunctionBuilder;
import java.util.ArrayList;
import java.util.List;
import picocli.CommandLine;

public class FunctionsConverter implements CommandLine.ITypeConverter<List<Function>> {

    @Override
    public List<Function> convert(String s) throws Exception {
        try {
            var functions = new ArrayList<Function>();
            for (var i : s.split(",")) {
                var func_item = i.split(":");
                var width = Double.parseDouble(func_item[1]);
                if (width <= 0) {
                    throw new CommandLine.TypeConversionException("Width should be positive!");
                }
                functions.add(new FunctionBuilder()
                        .byName(func_item[0])
                        .withWeight(width)
                        .build());
            }
            return functions;
        } catch (ArrayIndexOutOfBoundsException exception) {
            throw new CommandLine.TypeConversionException(exception.getMessage());
        }
    }
}
