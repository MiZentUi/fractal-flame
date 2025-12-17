package academy.cli.converters;

import academy.model.functions.Function;
import academy.utils.FunctionBuilder;
import picocli.CommandLine;
import java.util.ArrayList;
import java.util.List;

public class FunctionsConverter implements CommandLine.ITypeConverter<List<Function>> {

    @Override
    public List<Function> convert(String s) throws Exception {
        try {
            var functions = new ArrayList<Function>();
            for (var i : s.split(",")) {
                var func_item = i.split(":");
                functions.add(new FunctionBuilder().byName(func_item[0]).withWeight(Double.parseDouble(func_item[1])).build());
            }
            return functions;
        } catch (ArrayIndexOutOfBoundsException exception) {
            throw new CommandLine.TypeConversionException(exception.getMessage());
        }
    }
}
