from os import mkdir
from os.path import exists, dirname, join
import jinja2
from dot_gen import get_entity_mm


def main(debug=False):

    this_folder = dirname(__file__)

    entity_mm = get_entity_mm(debug)

    # Build Person model from person.ent file
    person_model = entity_mm.model_from_file(join(this_folder, 'invGato.ent'))

    def angular(s):
        return {
                'integer': 'int',
                'string': 'string'
        }.get(s.name, s.name)

    # Create output folder
    srcgen_folder = join(this_folder, 'srcgen')
    if not exists(srcgen_folder):
        mkdir(srcgen_folder)

    # Initialize template engine.
    jinja_env = jinja2.Environment(
        loader=jinja2.FileSystemLoader(this_folder),
        trim_blocks=True,
        lstrip_blocks=True)

    # Load Java template
    template = jinja_env.get_template('/templates/ver.template')

    for entity in person_model.entities:
        # For each entity generate java file
        with open(join(srcgen_folder,
                       "ver%s.html" % entity.name.capitalize()), 'w') as f:
            f.write(template.render(entity=entity))

    # Load Java template
    template = jinja_env.get_template('/templates/verCtrl.template')

    for entity in person_model.entities:
        # For each entity generate java file
        with open(join(srcgen_folder,
                       "ver%sCtrl.js" % entity.name.capitalize()), 'w') as f:
            f.write(template.render(entity=entity))

if __name__ == "__main__":
    main()
