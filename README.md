# Behavior3Gen

A secondary development tool for the Behavior3 low-code platform rule engine, designed to simplify the customization and
generation of .b3 JSON files for editor integration.

## Feature

- Generate `.b3` JSON files for custom nodes, which must be registered to the editor.
- Include node-level descriptions (as comments), instead of property-level details.
- Reflect properties in node names as parameters.
- Define custom default property values for new nodes.
- Automatically classify nodes based on their type (e.g., Action, Condition, Composite, or Decorator).
- Customize the node version for the Behavior3 architecture.

## Getting start

Install the package via:

```bash
go get github.com/Jacob953/behavior3gen@latest
```

## How to use

Follow these recommended steps to customize nodes:

1. Define a custom node with a specific type and property fields.
2. Implement an initialization function for the node.
3. Register the node in a global map.
4. Generate the file using `behavior3gen`.

For a detailed example, refer to the `/examples` directory.

## Authors

[@Jacob953](https://github.com/Jacob953)

## License

This repository is licensed under the [MIT License](LICENSE).