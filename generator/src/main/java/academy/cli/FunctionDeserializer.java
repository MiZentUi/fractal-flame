package academy.cli;

import academy.model.functions.Function;
import academy.utils.FunctionBuilder;
import com.fasterxml.jackson.core.JacksonException;
import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JavaType;
import com.fasterxml.jackson.databind.deser.std.StdDeserializer;
import com.fasterxml.jackson.databind.node.DoubleNode;
import com.fasterxml.jackson.databind.node.TextNode;
import java.io.IOException;

public class FunctionDeserializer extends StdDeserializer<Function> {

    public FunctionDeserializer() {
        super(Function.class);
    }

    protected FunctionDeserializer(Class<?> vc) {
        super(vc);
    }

    protected FunctionDeserializer(JavaType valueType) {
        super(valueType);
    }

    protected FunctionDeserializer(StdDeserializer<?> src) {
        super(src);
    }

    @Override
    public Function deserialize(JsonParser jsonParser, DeserializationContext deserializationContext)
            throws IOException, JacksonException {
        var node = jsonParser.getCodec().readTree(jsonParser);
        String name = ((TextNode) node.get("name")).asText();
        double weight = ((DoubleNode) node.get("weight")).doubleValue();
        return new FunctionBuilder().byName(name).withWeight(weight).build();
    }
}
