(jobs)=

# Jobs

```{eval-rst}
.. meta::
   :description: This page contains information about how to run and schedule jobs in Orchest.
```

Jobs are a way to schedule one-off or recurring {term}`pipelines <(data science) pipeline>` runs in Orchest.

A job can run multiple iterations of the same Pipeline over time
or small variations by using different parameters as inputs.
For example, you could create a job which uses the same ETL pipeline
but extracts data from a different data source for each Pipeline run.

```{figure} ../img/jobs-list.png
:align: center
:width: 768
:alt: List of jobs for a given Orchest project

The list of jobs for a given Orchest project.
```

Jobs take a snapshot of your project directory when they are created. Each of the job’s pipeline
runs copy the project directory snapshot and execute the files without changing the original
snapshot. This means jobs run consistently throughout their entire lifetime.

Different Pipeline runs that are part of the same job are completely isolated from a scheduling
perspective and do not affect each others' state.

```{note}
💡 Write data and large artifacts to the `/data` directory. This helps to save disk space by
not including them in the project directory snapshot. Alternatively, you can add the artifacts to
`.gitignore` as the ignored patterns are not copied to the snapshot.
```

```{tip}
👉 Check out our short video tutorials:

- [Adding parameters to a pipeline](https://app.tella.tv/story/cknrahyn9000409kyf4s2d3xm)
- [Running a pipeline as a job](https://app.tella.tv/story/cknr9nq1u000609kz9h0advvk)
```

(running-a-job)=

## Running a job in Orchest

To create and run a job in Orchest, follow these instructions:

1. Click on _Jobs_ in the left menu pane.
2. Click the _+ Create job_ button to configure your job.
3. Choose a _Job name_ and the _Pipeline_ you want to run the job for.
4. By default, the job is configured to run once immediately (_Now_).
   Alternatively, use the radio buttons to either schedule it for _Later_, or set a _Recurring_ schedule.
5. Press _Run job_.

```{figure} ../img/job-new.png
:align: center
:width: 428
:alt: Dialog that shows how to create a new job in Orchest

New job dialog.
```

To inspect the result of your job; click on the job you just created, choose a specific Pipeline run
(the one you want to inspect) and click on _View pipeline_. The Pipeline is now opened in
{term}`read-only mode` giving you the opportunity to check the logs or to open the
HTML version of you notebooks.

```{note}
💡 Upon job creation, Orchest (under the hood) takes a snapshot of the required environments.
This way you can freely iterate on and update your existing environments without worrying about
breaking your existing jobs.
```

(parametrize-pipeline-section)=

## Parametrizing pipelines and steps

Jobs run a specific Pipeline for a given set of parameters. If you define multiple values for the
same parameter, then the job will run the Pipeline once for every combination of parameter values.
You can think of job parameters as a [grid search](https://scikit-learn.org/stable/modules/grid_search.html),
i.e. looping over all combinations of values for different parameters.

```{note}
💡 Unlike {ref}`environment variables <environment-variables>`, you can define
Pipeline and step level parameters with the same name without one (automatically) overwriting
the other, you can access both values.
```

You can define Pipeline parameters at two levels:

- Pipelines: The parameters and their values will be available across every Pipeline step.
- Pipeline steps: The parameters will only be available in which they are defined.

### Editing pipeline parameters

1. Open a Pipeline via the _Pipelines_ option in the left menu pane.
2. Click on _SETTINGS_ in the top right corner.
3. Towards the top you will find the _Pipeline parameters_ section.
4. Input some JSON like {code}`{"my-param": <param-value>}`.
5. Make sure to _Save_ at the bottom of your screen.

### Editing pipeline step parameters

1. Open a Pipeline via the _Pipelines_ option in the left menu pane.
2. Click on a Pipeline step to open its _Properties_.
3. Towards the bottom you will find the _Parameters_ section.
4. Input some JSON like {code}`{"my-param": <param-value>}`.

(jobs-parameters)=

### Interacting with parameters through code

After you have set parameters for your Pipeline and/or steps you can use their values inside your
scripts (see {ref}`parameters API reference <api parameters>`).

Let's say you have set the following parameters on your Pipeline:

```json
{
  "vegetable": "carrot",
  "fruit": "banana"
}
```

And for your Pipeline step:

```json
{
  "candy": "chocolate",
  "fruit": "apple"
}
```

Then inside the Pipeline step you can access the parameters as follows:

```python
import orchest

# Get the parameters of the current step and the pipeline.
fruit = orchest.get_step_param("fruit")               # "apple"
vegetable = orchest.get_pipeline_param("vegetable")   # "carrot"
```

### Specify job parameters with a file

You can easily run a Pipeline for multiple parameter configurations by creating
a job parameters file.

If you place the file in the same folder as your Pipeline file the job parameters file will automatically be detected when creating a job.

For a Pipeline called `main.orchest` the job parameters file should be named `main.parameters.json`, and be put in the same folder as the pipeline file (both in the project directory).

You can also select a file manually when creating a job.

The JSON file should be formatted as below. **Note that wrapping the values in a list is required, even if you're assigning just one parameter value to a key.** It is allowed to omit keys you don't want to specify.

```json
{
  "pipeline_parameters": {
    "some_key": ["a", "list", "of", "values"]
  },
  "62a62810-336c-44c4-af6a-35228e8f2028": {
    "some_key": [1, 2, 3],
    "another_key": [1]
  }
}
```

You can find the step UUIDs in the pipeline file (e.g. `main.orchest`), pipelines are regular JSON files.

(running-a-parametrized-job)=

### Running a parametrized job

The procedure to run a parametrized job is very similar to running a job without any parameters.
Once you have followed any of the procedures above to parametrize your Pipeline:

1. Make sure you have defined some parameters or you will only be able to schedule the Pipeline as is.
2. Click on _Jobs_ in the left menu pane.
3. Click the _+ Create job_ button to configure your job.
4. Choose a _Job name_ and the _Pipeline_ you want to run the job for.
5. Your default set of parameters are pre-loaded. By clicking on the values a JSON editor opens,
   allowing you to add additional values you would like the Pipeline to run for.
6. If you would like to schedule the job to run at a specific time have a look at _Scheduling_. In
   case you don't want your job to run every combination of your parameter values, you can
   deselect them through the _Pipeline runs_ option.
7. Press _Run job_.
